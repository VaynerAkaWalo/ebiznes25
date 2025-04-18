plugins {
    id("application")
}

dependencies {
    implementation("org.xerial:sqlite-jdbc:3.49.1.0")
}

java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(8)
    }
}

application {
    mainClass = "Main"
}
