using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc.ModelBinding.Validation;

namespace Backend.Models;

public class DiscordMessage
{
    public string Id { get; set; } = null!;
    public string AuthorId { get; set; } = null!;
    public string ChannelId { get; set; } = null!;
    public string Content { get; set; } = null!;
    
    public DateTimeOffset Timestamp { get; set; }
    
    [JsonIgnore]
    [ValidateNever]
    public DiscordUser Author { get; set; } = null!;
    [JsonIgnore]
    [ValidateNever]
    public DiscordChannel Channel { get; set; } = null!;
}