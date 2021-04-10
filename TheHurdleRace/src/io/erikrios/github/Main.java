package io.erikrios.github;

public class Main {

    public static void main(String[] args) {
        int characterHeight = 7;
        int[] eachHurdleHeights = {2, 5, 4, 5, 2};
        System.out.println(hurdleRace(characterHeight, eachHurdleHeights));
    }

    private static int hurdleRace(int k, int[] height) {
        int dose = 0;
        for (int singleHeight : height) {
            if (k < singleHeight) {
                int heightRemainsRequired = singleHeight - k;
                k += heightRemainsRequired;
                dose += heightRemainsRequired;
            }
        }
        return dose;
    }
}
