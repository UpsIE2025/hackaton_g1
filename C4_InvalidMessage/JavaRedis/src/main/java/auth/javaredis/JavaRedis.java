/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */
package auth.javaredis;

/**
 *
 * @author amand
 */
public class JavaRedis {

    public static void main(String[] args) {
        RedisService redisService = new RedisService();
        MessageProcessor processor = new MessageProcessor(redisService);

        // Simulación de mensajes
        String validMessage = "{\"id\":1, \"sensor\":\"fuel\", \"value\":50}";
        String invalidMessage = "{\"id\":null, \"sensor\":\"\", \"value\":null}";

        processor.process(validMessage);
        processor.process(invalidMessage);
    }
}
