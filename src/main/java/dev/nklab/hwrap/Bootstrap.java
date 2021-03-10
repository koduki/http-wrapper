package dev.nklab.hwrap;

import java.util.logging.Logger;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.event.Observes;

import org.eclipse.microprofile.config.inject.ConfigProperty;

import io.quarkus.runtime.StartupEvent;

/**
 *
 * @author koduki
 */
@ApplicationScoped
public class Bootstrap {
    private static final Logger LOGGER = Logger.getLogger("htwrp");

    @ConfigProperty(name = "hwrap.cmd")
    String cmd;
    
    void onStart(@Observes StartupEvent event) {     
        LOGGER.info("cmd: " + cmd.replaceAll(",", " "));
    }
}