package io.erikrios.github;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {
        System.out.println(Arrays.toString(permutationEquation(new int[]{2, 3, 1})));
    }

    private static int[] permutationEquation(int[] p) {
        int[] results = new int[p.length];

        Map<Integer, Integer> hashMap = new HashMap<>();
        for (int i = 1; i <= p.length; i++) {
            hashMap.put(p[i - 1], i);
        }

        for (int i = 1; i <= p.length; i++) {
            results[i - 1] = hashMap.get(hashMap.get(i));
        }

        return results;
    }
}
