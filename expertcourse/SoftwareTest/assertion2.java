//https://recursionist.io/dashboard/problems/524
class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static int applyDiscount(int originalPrice, double discount) {
        double discountedPrice = originalPrice * (1.0 - discount);
        System.out.println((int) discountedPrice);

        return (int) discountedPrice;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
    }
}