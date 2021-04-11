package io.erikrios.github;

public class Main {

    public static void main(String[] args) {
        System.out.println(beautifulDays(20, 23, 6));
    }

    private static int beautifulDays(int i, int j, int k) {
        int discreteNumber = 0;

        for (int n = i; n <= j; n++) {
            StringBuilder numberString = new StringBuilder(String.valueOf(n));
            String reversedString = numberString.reverse().toString();
            reversedString = reversedString.replaceFirst("^0+(?!$)", "");

            double reversedDouble = Double.parseDouble(reversedString);

            double result = Math.abs((double) n - reversedDouble) / (double) k;
            if (result % 1.0 == 0)
                discreteNumber++;
        }
        return discreteNumber;
    }
}
