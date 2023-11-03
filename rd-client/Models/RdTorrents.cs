using System.Text.Json.Serialization;

namespace rd_client.Models;

public class RdTorrents
{
    [JsonPropertyName("id")]
    public string Id { get; set; }
    [JsonPropertyName("filename")]
    public string FileName { get; set; }
    [JsonPropertyName("hash")]
    public string Hash { get; set; }
    [JsonPropertyName("bytes")]
    public int Bytes { get; set; }
    [JsonPropertyName("host")]
    public string Host { get; set; }
    [JsonPropertyName("split")]
    public int Split { get; set; }
    [JsonPropertyName("progress")]
    public int Progress { get; set; }
    [JsonPropertyName("status")]
    public string Status { get; set; }
    [JsonPropertyName("added")]
    public string Added { get; set; }
    [JsonPropertyName("links")]
    public List<string> Link { get; set; }
    [JsonPropertyName("ended")]
    public string Ended { get; set; }
    [JsonPropertyName("speed")]
    public int? Speed { get; set; }
    [JsonPropertyName("seeders")]
    public int? Seeders { get; set; }
}
