import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int length = scanner.nextInt();
        int[] numbers = new int[length];

        for (int i = 0; i < length; i++) {
            numbers[i] = scanner.nextInt();
        }

        int count = 0;
        int sum = 0;

        for (int j = 0; j < length; j++) {
            for (int k = j; k < length; k++) {
                sum += numbers[k];
                if (sum < 0) {
                    count++;
                }
            }
            sum = 0;
        }

        System.out.print(count);
    }
}