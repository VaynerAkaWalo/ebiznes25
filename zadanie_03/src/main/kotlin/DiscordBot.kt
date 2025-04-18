package ebiznes

import net.dv8tion.jda.api.JDA
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.entities.Activity
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.EventListener
import net.dv8tion.jda.api.hooks.ListenerAdapter

class DiscordBot(private var token: String) : ListenerAdapter() {
    private lateinit var bot: JDA

    fun startBot() {
        bot = JDABuilder.createDefault(token)
            .setActivity(Activity.playing("Ebiznes"))
            .addEventListeners(this)
            .build()
            .awaitReady()
    }

    fun sendMessage() {
        val chanels = bot.getTextChannelsByName("ogólny", true)
        for(chan in chanels) {
            chan.sendMessage("Hello world").queue()
        }
    }

    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.message.author.isBot) {
            return
        }
        event.message.mentions

        if (event.message.mentions.isMentioned(bot.selfUser)) {
            event.channel.sendMessage("Odebralem wiadomosc " + event.message.contentRaw).queue()
        }
    }
}
