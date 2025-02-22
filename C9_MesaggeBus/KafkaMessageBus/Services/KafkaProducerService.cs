using System;
using System.Text.Json;
using System.Threading.Tasks;
using Confluent.Kafka;
using Microsoft.Extensions.Configuration;

public class KafkaProducerService
{
    private readonly string _bootstrapServers;
    private readonly string _topic;

    public KafkaProducerService(IConfiguration configuration)
    {      
        _bootstrapServers = configuration["Kafka:BootstrapServers"] ?? throw new ArgumentNullException("BootstrapServers is required.");
        _topic = configuration["Kafka:Topic"] ?? throw new ArgumentNullException("Topic is required.");
    }

    public async Task SendMessageAsync<T>(T message)
    {
        var config = new ProducerConfig { BootstrapServers = _bootstrapServers };

        using var producer = new ProducerBuilder<Null, string>(config).Build();
        var jsonMessage = JsonSerializer.Serialize(message);

        await producer.ProduceAsync(_topic, new Message<Null, string> { Value = jsonMessage });

        Console.WriteLine($"[Kafka Producer] Mensaje enviado: {jsonMessage}");
    }
}
