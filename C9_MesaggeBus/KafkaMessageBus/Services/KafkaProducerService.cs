using System;
using System.Text.Json;
using System.Threading.Tasks;
using Confluent.Kafka;
using Microsoft.Extensions.Configuration;

public class KafkaProducerService
{
    private readonly string _bootstrapServers;
    private readonly string _topic;
    private readonly IProducer<Null, string> _producer;

    public KafkaProducerService(IConfiguration configuration)
    {
        _bootstrapServers = configuration["Kafka:BootstrapServers"] ?? throw new ArgumentNullException("BootstrapServers is required.");
        _topic = configuration["Kafka:Topic"] ?? throw new ArgumentNullException("Topic is required.");

        var config = new ProducerConfig { BootstrapServers = _bootstrapServers };
        _producer = new ProducerBuilder<Null, string>(config).Build();
    }

    public async Task SendMessageAsync<T>(T message)
    {
        var jsonMessage = JsonSerializer.Serialize(message);
        await _producer.ProduceAsync(_topic, new Message<Null, string> { Value = jsonMessage });
        Console.WriteLine($"[Kafka Producer] Mensaje enviado: {jsonMessage}");
    }

    public void Dispose()
    {
        _producer.Flush();
        _producer.Dispose();
    }
}
