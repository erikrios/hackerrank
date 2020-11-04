//Complete this code or write your own from scratch

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

class Solution {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        in.nextLine();

        Map<String, Integer> phoneBook = new HashMap<>();

        for (int i = 0; i < n; i++) {
            String name = in.nextLine();
            int phone = in.nextInt();
            in.nextLine();
            phoneBook.put(name, phone);
        }
        while (in.hasNext()) {
            String s = in.nextLine();
            Integer phoneNumber = phoneBook.get(s);
            if (phoneNumber == null) {
                System.out.println("Not found");
            } else {
                System.out.println(s + "=" + phoneNumber);
            }
        }
    }
}