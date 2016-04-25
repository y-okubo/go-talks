// javac -cp jna.jar Main.java;java -cp .:jna.jar Main

import com.sun.jna.Library;
import com.sun.jna.Native;

interface HelloLib extends Library {
  HelloLib INSTANCE = (HelloLib) Native.loadLibrary("gofib", HelloLib.class);
  int fib(int n);
}

public class Main {
  public static void main(String[] args){
    HelloLib hello = HelloLib.INSTANCE;
    System.out.println(hello.fib(10));
  }
}
