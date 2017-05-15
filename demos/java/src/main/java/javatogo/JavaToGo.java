package javatogo;

import java.io.File;
import java.util.List;
import java.util.Arrays;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;
import com.sun.jna.Structure;

class GoString extends Structure {
  public String p;
  public long n;
  protected List<String> getFieldOrder() {
      return Arrays.asList(new String[] { "p", "n" });
  }
  public GoString(String x, long y) {
    p = x;
    n = y;
  }
}

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
    long PrintGoStr(GoString x);
    long PrintInt(long x);
    long PrintCStr(String x);
    long PrintStruct(String x);
    long PrintStructMeth(String x);
    String GetFirstJSONElement(String x);
  }

  public static void main(String[] args) {
    MY_LIB.PrintInt(42);
    MY_LIB.PrintCStr("Hi to C from Java");
    String pth = System.getProperty("user.dir") + "/../testdata.json";
    System.out.println(MY_LIB.GetFirstJSONElement(pth));

    // These aren't working properly
    GoString gs = new GoString("hi", 2L);
    MY_LIB.PrintGoStr(gs);
    MY_LIB.PrintStruct("42");
    MY_LIB.PrintStructMeth("42");
  }
}
