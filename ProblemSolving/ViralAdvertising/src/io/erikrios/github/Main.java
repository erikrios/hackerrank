package io.erikrios.github;

public class Main {

    public static void main(String[] args) {
        System.out.println(viralAdvertising(5));
    }

    private static int viralAdvertising(int n) {
        int cumulative = 0;
        int shared = 5;
        int liked = 0;

        for (int i = 0; i < n; i++) {
            liked = (int) Math.floor((double) shared / 2.0);
            cumulative += liked;
            shared = liked * 3;
        }

        return cumulative;
    }
}
