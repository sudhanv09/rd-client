using System.Text.Json.Serialization;

namespace rd_client.Models;

public class RdMagnetResult
{
    [JsonPropertyName("id")]
    public string Id { get; set; }
    [JsonPropertyName("uri")]
    public string Uri { get; set; }
}