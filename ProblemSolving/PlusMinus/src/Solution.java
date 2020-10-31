import java.util.Scanner;

public class Solution {

    // Complete the plusMinus function below.
    static void plusMinus(int[] arr) {
        int plusCount = 0;
        int minusCount = 0;
        int zeroCount = 0;

        for (int value : arr)
            if (value > 0)
                plusCount++;
            else if (value < 0)
                minusCount++;
            else
                zeroCount++;

        double plusProportion = (double) plusCount / (double) arr.length;
        double minusProportion = (double) minusCount / (double) arr.length;
        double zeroProportion = (double) zeroCount / (double) arr.length;

        System.out.printf("%6f\n", plusProportion);
        System.out.printf("%6f\n", minusProportion);
        System.out.printf("%6f", zeroProportion);

    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int[] arr = new int[n];

        String[] arrItems = scanner.nextLine().split(" ");
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int i = 0; i < n; i++) {
            int arrItem = Integer.parseInt(arrItems[i]);
            arr[i] = arrItem;
        }

        plusMinus(arr);

        scanner.close();
    }
}
