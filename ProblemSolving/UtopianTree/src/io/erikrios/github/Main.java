package io.erikrios.github;

public class Main {

    public static void main(String[] args) {
        System.out.println(utopianTree(4));
    }

    private static int utopianTree(int n) {
        int tall = 0;
        for (int i = 0; i <= n; i++) {
            if (i % 2 == 0)
                tall++;
            else tall *= 2;
        }
        return tall;
    }
}
