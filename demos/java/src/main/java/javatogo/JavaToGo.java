package javatogo;

import java.io.File;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;

public class JavaToGo {
  static MyLib MY_LIB;
  static {
    String os = System.getProperty("os.name").toLowerCase();
    String libExtension;
    if (os.contains("mac os")) {
      libExtension = "dylib";
    } else if (os.contains("windows")) {
      libExtension = "dll";
    } else {
      libExtension = "so";
    }

    String pwd = System.getProperty("user.dir");
    String lib = pwd + "/mylib." + libExtension;
    MY_LIB = (MyLib) Native.loadLibrary(lib, MyLib.class);
  }

  public interface MyLib extends Library {
    long PrintInt(long x);
    long PrintCStr(String x);
  }

  public static void main(String[] args) {
    System.out.println("Java says: about to call Go ..");
    MY_LIB.PrintCStr("Hi to C from Java");
    MY_LIB.PrintInt(42);
  }
}
