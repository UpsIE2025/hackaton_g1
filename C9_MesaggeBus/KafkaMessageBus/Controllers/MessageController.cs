using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

[Route("api/messages")]
[ApiController]
public class MessageController : ControllerBase
{
    private readonly KafkaProducerService _kafkaProducer;

    public MessageController(KafkaProducerService kafkaProducer)
    {
        _kafkaProducer = kafkaProducer;
    }

    [HttpPost]
    public async Task<IActionResult> SendMessage([FromBody] MarketingEvent marketingEvent)
    {
        await _kafkaProducer.SendMessageAsync(marketingEvent);
        return Ok("Mensaje enviado a Kafka.");
    }
}

public class MarketingEvent
{
    public string Campaign { get; set; } = string.Empty;
    public string Action { get; set; } = string.Empty;
    public string User { get; set; } = string.Empty;
}
