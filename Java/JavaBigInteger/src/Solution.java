import java.math.BigInteger;
import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        Scanner scanner = new Scanner(System.in);
        BigInteger firstNumber = scanner.nextBigInteger();
        BigInteger secondNumber = scanner.nextBigInteger();

        System.out.println(firstNumber.add(secondNumber));
        System.out.print(firstNumber.multiply(secondNumber));
    }
}