package io.erikrios.github;

import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        int[] heights = {1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7};
        String word = "zaba";
        System.out.println(designerPdfViewer(heights, word));
    }

    private static int designerPdfViewer(int[] h, String word) {
        String allCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".toLowerCase();
        word = word.toLowerCase();

        List<Integer> indexs = new ArrayList<>();

        for (int i = 0; i < word.length(); i++) {
            int index = allCharacters.indexOf(word.charAt(i));
            indexs.add(index);
        }

        int highest = 0;

        for (int index : indexs) {
            int height = h[index];
            if (height > highest)
                highest = height;
        }

        return highest * word.length();
    }
}
