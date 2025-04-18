package ebiznes

import net.dv8tion.jda.api.JDA
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.entities.Activity

class DiscordBot(private var token: String) {
    private lateinit var bot: JDA

    fun startBot() {
        bot = JDABuilder.createDefault(token)
            .setActivity(Activity.playing("Ebiznes"))
            .build()
            .awaitReady()
    }

    fun sendMessage() {
        val chanels = bot.getTextChannelsByName("ogólny", true)
        for(chan in chanels) {
            chan.sendMessage("Hello world").queue()
        }
    }
}
