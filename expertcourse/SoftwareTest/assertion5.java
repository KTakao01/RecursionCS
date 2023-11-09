//https://recursionist.io/dashboard/problems/527
import java.util.Arrays;

class Assertion {
    public static void run(boolean b) throws Exception {
        if (!b) throw new Exception("Assertion Error");
        else System.out.println("Passed the assertion test.");
    }
}

class MyClass {
    public static int primesUpToNCount(int n) {
        boolean[] cache = new boolean[n];
        Arrays.fill(cache, true);
        for (int i = 2; i <= Math.floor(Math.sqrt(n)); i++) {
            if (!cache[i]) continue;
            for (int p = i * 2; p < n; p += i) {
                cache[p] = false;
            }
        }

        int primeCount = 0;
        for (int i = 2; i < cache.length; i++) {
            if (cache[i]) primeCount += 1;
        }
        //System.out.println(primeCount);
        return primeCount;
    }
    
    public static int primesUpToNCountBF(int n) {
        int counter = 0;
        for(int i = 2; i < n; i++) {
            if(isPrime(i)) counter ++;
        }
        //System.out.println(counter);
        return counter;
    }

    public static boolean isPrime(int n) {
        if(n < 2) return false;
        for(int i = 2; i * i <= n; i++) {
            if(n % i == 0) return false;
        }
        return true;
    }

    public static void main(String[] args) throws Exception {
        // ここから書きましょう
        Assertion.run(primesUpToNCount(2) == primesUpToNCountBF(2));
        Assertion.run(primesUpToNCount(3) == primesUpToNCountBF(3));
        Assertion.run(primesUpToNCount(5) == primesUpToNCountBF(5));
        Assertion.run(primesUpToNCount(10) == primesUpToNCountBF(10));
        Assertion.run(primesUpToNCount(19) == primesUpToNCountBF(19));
        Assertion.run(primesUpToNCount(25) == primesUpToNCountBF(25));
        Assertion.run(primesUpToNCount(1000) == primesUpToNCountBF(1000));
        //primesUpToNCountBF(2);
        //primesUpToNCountBF(3);
        //primesUpToNCountBF(5);        
        //primesUpToNCountBF(10);
        //primesUpToNCountBF(19);
        //primesUpToNCountBF(25);
        //primesUpToNCountBF(1000);
        //primesUpToNCount(2);
        //primesUpToNCount(3);
        //primesUpToNCount(5);
        //primesUpToNCount(10);
        //primesUpToNCount(19);
        //primesUpToNCount(25);
        //primesUpToNCount(1000);
    }
}