package ebiznes

fun main(args: Array<String>) {
    val bot = DiscordBot(System.getenv("DC_TOKEN"))
    bot.startBot()

    bot.sendMessage()
}
