package auth.javaredis;

import redis.clients.jedis.Jedis;

/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */

/**
 *
 * @author amand
 */
public class RedisService {
    private static final String REDIS_HOST = "http://127.0.0.1:6379";
    private Jedis jedis;

    public RedisService() {
        this.jedis = new Jedis(REDIS_HOST);
        System.out.println("✅ Conectado a Redis en " + REDIS_HOST);
    }

    public void saveValidMessage(String message) {
        jedis.lpush("valid_messages", message);
        System.out.println("✅ Mensaje válido almacenado en Redis: " + message);
    }

    public void saveInvalidMessage(String message) {
        jedis.lpush("invalid_messages", message);
        System.out.println("❌ Mensaje inválido almacenado en Redis: " + message);
    }
}
