using System;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using Confluent.Kafka;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using StackExchange.Redis;

public class KafkaConsumerService : BackgroundService
{
    private readonly string _bootstrapServers;
    private readonly string _topic;
    private readonly IConsumer<Ignore, string> _consumer;

    public KafkaConsumerService(IConfiguration configuration)
    {
        _bootstrapServers = configuration["Kafka:BootstrapServers"] ?? "localhost:9092";
        _topic = configuration["Kafka:Topic"] ?? "default_topic";

        var config = new ConsumerConfig
        {
            BootstrapServers = _bootstrapServers,
            GroupId = "test-group",
            AutoOffsetReset = AutoOffsetReset.Earliest
        };

        _consumer = new ConsumerBuilder<Ignore, string>(config).Build();
        _consumer.Subscribe(_topic);
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        var config = new ConsumerConfig
    {
        BootstrapServers = _bootstrapServers,
        GroupId = "marketing-consumer-group",
        AutoOffsetReset = AutoOffsetReset.Earliest
    };

    using var consumer = new ConsumerBuilder<Ignore, string>(config).Build();
    consumer.Subscribe(_topic);

    while (!stoppingToken.IsCancellationRequested)
    {
        try
        {
            var consumeResult = await Task.Run(() => consumer.Consume(stoppingToken), stoppingToken); // Usamos Task.Run para hacer la operación asincrónica
            var message = consumeResult.Message.Value;

            Console.WriteLine($"[Kafka Consumer] Mensaje recibido: {message}");

            // Guardar en Redis
            await _redisDb.ListLeftPushAsync("marketing-events", message);
        }
        catch (ConsumeException ex)
        {
            Console.WriteLine($"[Kafka Consumer] Error: {ex.Error.Reason}");
        }
    }
    }
}
