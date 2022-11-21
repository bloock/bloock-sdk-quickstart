package com.bloock.quickstart.utils;

public class Logger {

    public static final String ANSI_RESET = "\u001B[0m";
    public static final String ANSI_YELLOW = "\u001B[33m";
    public static final String ANSI_RED = "\u001B[31m";
    public static final String ANSI_GREEN = "\u001B[32m";
    public static final String ANSI_CYAN = "\u001B[36m";

    public static void success(String message) {
        System.out.println(ANSI_GREEN + message + ANSI_RESET);
    }

    public static void warn(String message) {
        System.out.println(ANSI_YELLOW + message + ANSI_RESET);
    }

    public static void info(String message) {
        System.out.println(ANSI_CYAN + message + ANSI_RESET);
    }

    public static void error(String message) {
        System.out.println(ANSI_RED + message + ANSI_RESET);
    }    
}
