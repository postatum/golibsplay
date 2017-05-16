package javatogo;

import java.util.List;
import java.util.Arrays;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;
import com.sun.jna.Structure;


public class JavaToGo {
  static MyLib MY_LIB;
  static {
    String pwd = System.getProperty("user.dir");
    String lib = pwd + "/mylib.so";
    MY_LIB = (MyLib) Native.loadLibrary(lib, MyLib.class);
  }

  public interface MyLib extends Library {
    long PrintGoStr(GoString.ByValue x);
    long PrintInt(long x);
    long PrintCStr(String x);
    long PrintStruct(String x);
    long PrintStructMeth(String x);
    String GetFirstJSONElement(String x);
  }

  public static class GoString extends Structure {
    public static class ByValue extends GoString implements Structure.ByValue {}
    public String p;
    public long n;
    protected List getFieldOrder() {
        return Arrays.asList(new String[] { "p", "n" });
    }
  }

  public static void main(String[] args) {
    System.out.println("\nPrintInt() will run now:");
    MY_LIB.PrintInt(42);

    System.out.println("\nPrintCStr() will run now:");
    MY_LIB.PrintCStr("Hi to C from Java");

    System.out.println("\nGetFirstJSONElement() will run now:");
    String pth = System.getProperty("user.dir") + "/../testdata.json";
    System.out.println(MY_LIB.GetFirstJSONElement(pth));

    System.out.println("\nPrintGoStr() will run now:");
    GoString.ByValue gs = new GoString.ByValue();
    gs.p = "Hello Java!";
    gs.n = 11;
    MY_LIB.PrintGoStr(gs);

    // These aren't working properly
    // MY_LIB.PrintStruct("42");
    // MY_LIB.PrintStructMeth("42");
  }
}
