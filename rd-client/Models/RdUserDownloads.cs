using System.Text.Json.Serialization;

namespace rd_client.Models;

public class RdUserDownloads
{
    [JsonPropertyName("id")]
    public string Id { get; set; }
    [JsonPropertyName("filename")]
    public string FileName { get; set; }
    [JsonPropertyName("mimeType")]
    public string MimeType { get; set; }
    [JsonPropertyName("filesize")]
    public int Filesize { get; set; }
    [JsonPropertyName("link")]
    public string Link { get; set; }
    [JsonPropertyName("host")]
    public string Host { get; set; }
    [JsonPropertyName("chunks")]
    public int Chunks { get; set; }
    [JsonPropertyName("download")]
    public string Download { get; set; }
    [JsonPropertyName("generated")]
    public string Generated { get; set; }
    [JsonPropertyName("type")]
    public string? Type { get; set; }
}