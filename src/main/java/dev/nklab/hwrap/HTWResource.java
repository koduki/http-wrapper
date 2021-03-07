package dev.nklab.hwrap;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

import org.eclipse.microprofile.config.inject.ConfigProperty;

@Path("/")
public class HTWResource {

    @ConfigProperty(name = "hwrap.cmd")
    String cmd;

    @GET
    @Produces(MediaType.TEXT_PLAIN)
    public Response exec(@QueryParam("args") String params) throws Exception {
        if (params == null) {
            params = "";
        }

        System.out.print(cmd.replaceAll(",", ""));
        System.out.println(" " + params.replaceAll(",", ""));

        var cmds = Arrays.stream(cmd.split(",")).map(x -> x.strip()).toArray(String[]::new);
        var args = Arrays.stream(params.split(",")).map(x -> x.strip()).toArray(String[]::new);

        var r = execCommand(concat(cmds, args));

        System.out.println(r.get("stdout"));

        String err = r.get("stderr");

        if (r.get("status").equals("0")) {
            return Response.ok("success").build();
        } else {
            System.err.println(err);
            return Response.status(Response.Status.INTERNAL_SERVER_ERROR).entity(err).build();
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

    private Map<String, String> execCommand(String... cmds) throws IOException, InterruptedException {
        var result = new HashMap<String, String>();
        var proc = Runtime.getRuntime().exec(cmds);

        result.put("stdout", toString(proc.getInputStream()));
        result.put("stderr", toString(proc.getErrorStream()));
        result.put("status", Integer.toString(proc.waitFor()));

        return result;
    }

    private String toString(InputStream input) throws IOException {
        try {
            return new BufferedReader(new InputStreamReader(input)).lines()
                    .collect(Collectors.joining(System.lineSeparator()));
        } finally {
            input.close();
        }
    }
}