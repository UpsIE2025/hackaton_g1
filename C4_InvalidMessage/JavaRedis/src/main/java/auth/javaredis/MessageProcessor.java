/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package auth.javaredis;

/**
 *
 * @author amand
 */
public class MessageProcessor {
    private RedisService redisService;

    public MessageProcessor(RedisService redisService) {
        this.redisService = redisService;
    }

    public void process(String message) {
        if (isInvalidMessage(message)) {
            redisService.saveInvalidMessage(message);
        } else {
            redisService.saveValidMessage(message);
        }
    }

    private boolean isInvalidMessage(String message) {
        return message.contains("\"id\":null") || message.contains("\"sensor\":\"\"") || message.contains("\"value\":null");
    }
}
