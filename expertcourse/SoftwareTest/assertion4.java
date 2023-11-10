//https://recursionist.io/dashboard/problems/526
import java.util.regex.Pattern;

class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static boolean isValidPassword(String password) {
        boolean case1 = !password.contains(" ");
        boolean case2 = password.length() >= 6;
        boolean case3 = Pattern.matches("([0-9])", password);
        boolean case4 = !password.equals(password.toUpperCase()) && !password.equals(password.toLowerCase());
        return case1 && case2 && case3 && case4;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
        Assertion.run(isValidPassword("Aa 12E4") == false);
        Assertion.run(isValidPassword("abe12cd") == false);
        Assertion.run(isValidPassword("Aa1") == false);
        Assertion.run(isValidPassword("AaBbC23") == false);
        Assertion.run(isValidPassword("123456") == false);
        Assertion.run(isValidPassword("uYjkmese") == false);
        Assertion.run(isValidPassword("tK 23bmd") == false);
        Assertion.run(isValidPassword("aaaaa1") == false);
        Assertion.run(isValidPassword("AAAAA1") == false);
    }
}