using System.Text.Json.Serialization;

namespace rd_client.Models;

public class RdUnrestrict
{
    [JsonPropertyName("id")]
    public string Id { get; set; }
    [JsonPropertyName("filename")]
    public string FileName { get; set; }
    [JsonPropertyName("mimeType")]
    public string MimeType { get; set; }
    [JsonPropertyName("filesize")]
    public Int64 Filesize { get; set; }
    [JsonPropertyName("link")]
    public string Link { get; set; }
    [JsonPropertyName("host")]
    public string Host { get; set; }
    [JsonPropertyName("chunks")]
    public Int64 Chunks { get; set; }
    [JsonPropertyName("crc")]
    public int Crc { get; set; }
    [JsonPropertyName("download")]
    public string Download { get; set; }
    [JsonPropertyName("streamable")]
    public int Streamable { get; set; }
}