package ebiznes

import net.dv8tion.jda.api.JDA
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.entities.Activity
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter

class DiscordBot(private var token: String) : ListenerAdapter() {
    private lateinit var bot: JDA
    private var categories: List<Category> = listOf(Category(1L, "słodycze"), Category(2L, "pieczywo"));

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

        if (event.message.mentions.isMentioned(bot.selfUser) && event.message.contentRaw.contains("kategorie")) {
            event.channel.sendMessage("Dostepne kategoriee to: ${categories.stream().map { c -> c.name }.toList()}").queue()
        }
    }
}
