import java.util.*;

public class Main {

	public static void main(String[] args) {
		int[] initialNumbers = {1, 2, 2, 3, 1, 2};

		List<Integer> numbers = new ArrayList<Integer>();

		for (int i = 0; i < initialNumbers.length; i++)
			numbers.add(initialNumbers[i]);

		System.out.println(pickingNumbers(numbers));
	}

	public static int pickingNumbers(List<Integer> a) {
		Collections.sort(a);
		
		Set<Integer> uniqueNumbers = new HashSet<Integer>();
		uniqueNumbers.addAll(a);

		List<List<Integer>> subNumbers = new ArrayList<List<Integer>>();

		Iterator<Integer> iterators = uniqueNumbers.iterator();

		while (iterators.hasNext()) {
			List<Integer> subs = new ArrayList<Integer>();
			int uniqueNumber = iterators.next();

			for (int  i = 0; i < a.size(); i++) {
				int number = a.get(i);
				if (number >= uniqueNumber && Math.abs(uniqueNumber - number) <= 1)
					subs.add(number);
			}

			subNumbers.add(subs);
		}

		int longest = 0;
		for (int i = 0; i < subNumbers.size(); i++) {
			int subLength = subNumbers.get(i).size();
			if (subLength > longest)
				longest = subLength;
		}

		return longest;
	}
}
