package dev.nklab.hwrap;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Consumer;
import java.util.logging.Logger;
import java.util.stream.Stream;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

import org.eclipse.microprofile.config.inject.ConfigProperty;

@Path("/")
public class HwrapResource {
    private static final Logger LOGGER = Logger.getLogger("htwrp");

    @ConfigProperty(name = "hwrap.cmd")
    String cmd;

    @GET
    @Produces(MediaType.TEXT_PLAIN)
    @Path("/healthcheck")
    public Response healthcheck() throws Exception {
        return Response.ok("cmd: " + cmd).build();
    }

    @GET
    @Produces(MediaType.TEXT_PLAIN)
    public Response exec(@QueryParam("args") String params) throws Exception {
        if (params == null) {
            params = "";
        }
        LOGGER.info(cmd.replaceAll(",", " ") + " " + params.replaceAll(",", " "));

        var cmds = Arrays.stream(cmd.split(",")).map(x -> x.strip()).toArray(String[]::new);
        var args = Arrays.stream(params.split(",")).map(x -> x.strip()).toArray(String[]::new);

        var r = execCommand(
            s -> System.out.println(s), 
            s -> System.err.println(s), 
            concat(cmds, args)
        );

        if (r.get("status").equals("0")) {
            return Response.ok("success").build();
        } else {
            return Response.status(Response.Status.INTERNAL_SERVER_ERROR)
                            .entity(r.get("stderr"))
                            .build();
        }
    }

    public static class CommandExecException extends RuntimeException {
        private static final long serialVersionUID = 1L;

        public CommandExecException(String message) {
            super(message);
        }
    }

    private String[] concat(String[] args1, String[] args2) {
        var xs = new String[args1.length + args2.length];
        System.arraycopy(args1, 0, xs, 0, args1.length);
        System.arraycopy(args2, 0, xs, args1.length, args2.length);
        return xs;
    }

    private Map<String, String> execCommand(Consumer<String> stdProc, Consumer<String> errProc, String... cmds)
            throws IOException, InterruptedException {
        var result = new HashMap<String, String>();
        var proc = Runtime.getRuntime().exec(cmds);

        result.put("stdout", read(proc.getInputStream(), (s) -> System.out.println(s)));
        result.put("stderr", read(proc.getErrorStream(), (s) -> System.err.println(s)));
        result.put("status", Integer.toString(proc.waitFor()));

        return result;
    }

    private String read(InputStream input, Consumer<String> callback) {
        var sb = new StringBuilder();
        try (Stream<String> s = new BufferedReader(new InputStreamReader(input)).lines()) { // 自動close
            s.map((l) -> {
                sb.append(l);
                sb.append(System.lineSeparator());
                return l;
            }).forEach(callback);
        }
        return sb.toString();
    }
}