import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int totalNumber = Integer.parseInt(scanner.nextLine());
        List<List<Integer>> numbers = new ArrayList<>();

        for (int i = 0; i < totalNumber; i++) {
            String input = scanner.nextLine();
            String[] arrInput = input.split(" ");

            List<Integer> listNumber = new ArrayList<>();

            int length = Integer.parseInt(arrInput[0]);

            for (int j = 0; j < length; j++) {
                listNumber.add(Integer.parseInt(arrInput[j + 1]));
            }

            numbers.add(listNumber);
        }

        int totalQuery = Integer.parseInt(scanner.nextLine());
        List<List<Integer>> queries = new ArrayList<>();

        for (int i = 0; i < totalQuery; i++) {

            String[] strQuery = scanner.nextLine().split(" ");

            List<Integer> listQuery = new ArrayList<>();

            for (String s : strQuery) {
                listQuery.add(Integer.parseInt(s));
            }

            queries.add(listQuery);
        }

        for (int i = 0; i < queries.size(); i++) {
            try {
                int numbersElementQuery = queries.get(i).get(0);
                int numberQuery = queries.get(i).get(1);
                int number = numbers.get(numbersElementQuery - 1).get(numberQuery - 1);
                System.out.println(number);
            } catch (IndexOutOfBoundsException e) {
                System.out.println("ERROR!");
            }
        }
    }
}
