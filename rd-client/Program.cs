
using System.Diagnostics;
using rd_client.Lib;

await foreach (string text in new ClipboardWatcher().PollClipboard())
{
    Console.WriteLine(text);
}




