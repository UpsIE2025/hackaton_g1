/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package auth.javaredis;

/**
 *
 * @author amand
 */
import redis.clients.jedis.Jedis;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import java.util.List;

@Path("/messages")
public class RedisController {

    private static final String REDIS_HOST = "http://127.0.0.1:6379";
    private Jedis jedis = new Jedis(REDIS_HOST);

    @POST
    @Path("/send")
    @Consumes(MediaType.APPLICATION_JSON)
    public String sendMessage(String message) {
        if (isInvalidMessage(message)) {
            jedis.lpush("invalid_messages", message);
            return "❌ Mensaje inválido almacenado en Redis.";
        } else {
            jedis.lpush("valid_messages", message);
            return "✅ Mensaje válido procesado y almacenado en Redis.";
        }
    }

    @GET
    @Path("/valid")
    @Produces(MediaType.APPLICATION_JSON)
    public List<String> getValidMessages() {
        return jedis.lrange("valid_messages", 0, -1);
    }

    @GET
    @Path("/invalid")
    @Produces(MediaType.APPLICATION_JSON)
    public List<String> getInvalidMessages() {
        return jedis.lrange("invalid_messages", 0, -1);
    }

    private boolean isInvalidMessage(String message) {
        return message.contains("\"id\":null") || message.contains("\"sensor\":\"\"") || message.contains("\"value\":null");
    }
}
