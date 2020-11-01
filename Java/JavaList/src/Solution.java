import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int totalNumber = Integer.parseInt(scanner.nextLine());

        String[] strNumbers = scanner.nextLine().split(" ");

        List<Integer> numbers = new ArrayList<>();

        for (int i = 0; i < totalNumber; i++)
            numbers.add(Integer.parseInt(strNumbers[i]));

        int totalQuery = Integer.parseInt(scanner.nextLine());

        while (totalQuery > 0) {
            String query = scanner.nextLine();
            if (query.equalsIgnoreCase("Insert")) {
                String[] strQueryNumbers = scanner.nextLine().split(" ");
                numbers.add(Integer.parseInt(strQueryNumbers[0]), Integer.parseInt(strQueryNumbers[1]));
            } else {
                int queryIndex = Integer.parseInt(scanner.nextLine());
                numbers.remove(queryIndex);
            }

            totalQuery--;
        }

        for (int number : numbers) {
            System.out.printf("%d ", number);
        }
    }
}
