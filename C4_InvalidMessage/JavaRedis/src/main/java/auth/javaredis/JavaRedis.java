/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 */
package auth.javaredis;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import redis.clients.jedis.Jedis;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.List;
import java.util.List;
import org.json.JSONObject;
import redis.clients.jedis.Jedis;

/**
 *
 * @author amand
 */
public class JavaRedis {

    private static final String BASE_URI = "http://localhost:8080/";

    /*public static void main(String[] args) {

        RedisService redisService = new RedisService();
        MessageProcessor processor = new MessageProcessor(redisService);

        // Simulación de mensajes
        String validMessage = "{\"id\":1, \"sensor\":\"fuel\", \"value\":50}";
        String invalidMessage = "{\"id\":null, \"sensor\":\"\", \"value\":null}";

        processor.process(validMessage);
        processor.process(invalidMessage);*
        
    }*/
    private static final String REDIS_HOST = "http://127.0.0.1:6379";
    private static Jedis jedis;

    public static void main(String[] args) throws IOException {
        jedis = new Jedis(REDIS_HOST);
        System.out.println("✅ Servidor REST en ejecución en http://localhost:8080/");

        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
        server.createContext("/messages/send", new SendMessageHandler());
        server.createContext("/messages/valid", new GetValidMessagesHandler());
        server.createContext("/messages/invalid", new GetInvalidMessagesHandler());
        server.setExecutor(null);
        server.start();
    }

    static class SendMessageHandler implements HttpHandler {

        @Override
        public void handle(HttpExchange exchange) throws IOException {
            if ("POST".equals(exchange.getRequestMethod())) {
                byte[] requestBody = exchange.getRequestBody().readAllBytes();
                String message = new String(requestBody);

                if (isInvalidMessage(message)) {
                    jedis.lpush("invalid_messages", message);
                    sendResponse(exchange, "❌ Mensaje inválido almacenado en Redis.", 200);
                } else {
                    jedis.lpush("valid_messages", message);
                    sendResponse(exchange, "✅ Mensaje válido almacenado en Redis.", 200);
                }
            } else {
                sendResponse(exchange, "Método no permitido", 405);
            }
        }
    }

    static class GetValidMessagesHandler implements HttpHandler {

        @Override
        public void handle(HttpExchange exchange) throws IOException {
            List<String> messages = jedis.lrange("valid_messages", 0, -1);
            sendResponse(exchange, messages.toString(), 200);
        }
    }

    static class GetInvalidMessagesHandler implements HttpHandler {

        @Override
        public void handle(HttpExchange exchange) throws IOException {
            List<String> messages = jedis.lrange("invalid_messages", 0, -1);
            sendResponse(exchange, messages.toString(), 200);
        }
    }

    private static boolean isInvalidMessage(String message) {
        try {
            JSONObject json = new JSONObject(message);

            // Verificar si existen las claves esperadas
            if (!json.has("id") || !json.has("sensor") || !json.has("value")) {
                return true;
            }

            // Validar que los valores no sean null o vacíos
            if (json.isNull("id") || json.get("id").toString().trim().isEmpty()) {
                return true;
            }
            if (json.isNull("sensor") || json.getString("sensor").trim().isEmpty()) {
                return true;
            }
            if (json.isNull("value") || json.get("value").toString().trim().isEmpty()) {
                return true;
            }

            return false; // Si todas las validaciones pasan, el mensaje es válido

        } catch (Exception e) {
            return true; // Si hay un error en el formato del JSON, se considera inválido
        }
    }

    private static void sendResponse(HttpExchange exchange, String response, int statusCode) throws IOException {
        exchange.sendResponseHeaders(statusCode, response.length());
        OutputStream os = exchange.getResponseBody();
        os.write(response.getBytes());
        os.close();
    }
}
