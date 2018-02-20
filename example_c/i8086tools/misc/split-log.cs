using System;
using System.IO;
using System.Text;

class Test {
	private static int no = 0;
	private static StreamWriter GetWrite() {
		var fn = string.Format("log-{0:0000}", no++);
		return new StreamWriter(fn, false, Encoding.Default);
	}

	public static int Main(string[] args) {
		if (args.Length != 1) return 1;
		var sr = new StreamReader(args[0], Encoding.Default);
		var sw = GetWrite();
		int ln = 0;
		string line;
		while ((line = sr.ReadLine()) != null) {
			sw.WriteLine(line);
			if (++ln == 200000) {
				sw.Close();
				sw = GetWrite();
				ln = 0;
			}
		}
		sw.Close();
		sr.Close();
		return 0;
	}
}
