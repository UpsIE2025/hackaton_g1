using System;
using System.Threading;
using System.Threading.Tasks;
using Confluent.Kafka;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using StackExchange.Redis;

public class KafkaConsumerService : BackgroundService
{
    private readonly IConsumer<Ignore, string> _consumer;
    private readonly IDatabase _redisDb;

    public KafkaConsumerService(IConfiguration configuration, IConnectionMultiplexer redis)
    {
        var bootstrapServers = configuration["Kafka:BootstrapServers"] ?? "localhost:9092";
        var topic = configuration["Kafka:Topic"] ?? "default_topic";

        _redisDb = redis.GetDatabase();

        var config = new ConsumerConfig
        {
            BootstrapServers = bootstrapServers,
            GroupId = "test-group",
            AutoOffsetReset = AutoOffsetReset.Earliest
        };

        _consumer = new ConsumerBuilder<Ignore, string>(config).Build();
        _consumer.Subscribe(topic);
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        while (!stoppingToken.IsCancellationRequested)
        {
            try
            {
                var consumeResult = _consumer.Consume(stoppingToken);
                var message = consumeResult.Message.Value;

                Console.WriteLine($"[Kafka Consumer] Mensaje recibido: {message}");

                await _redisDb.ListLeftPushAsync("marketing-events", message);
            }
            catch (ConsumeException ex)
            {
                Console.WriteLine($"[Kafka Consumer] Error: {ex.Error.Reason}");
            }
        }
    }

    public override void Dispose()
    {
        _consumer.Close();
        _consumer.Dispose();
        base.Dispose();
    }
}
