package io.erikrios.github;

public class Main {

    public static void main(String[] args) {
        System.out.println(angryProfessor(2, new int[]{0, -1, 2, 1}));
    }

    private static String angryProfessor(int k, int[] a) {
        int onTimeStudent = 0;
        for (int arrival : a) {
            if (arrival <= 0)
                onTimeStudent++;
        }
        System.out.println(onTimeStudent);
        if (onTimeStudent < k) return "YES";
        else return "NO";
    }
}
