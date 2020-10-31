import java.math.BigInteger;
import java.util.Scanner;

public class Solution {

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        String n = scanner.nextLine();

        BigInteger number = new BigInteger(n);
        boolean isPrime = number.isProbablePrime(1);

        if (isPrime)
            System.out.print("prime");
        else
            System.out.print("not prime");

        scanner.close();
    }
}
