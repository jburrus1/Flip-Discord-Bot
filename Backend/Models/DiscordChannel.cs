using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc.ModelBinding.Validation;

namespace Backend.Models;

public class DiscordChannel
{
    public string Id { get; set; } = null!;
    public string Name { get; set; } = null!;
    public string ServerId { get; set; } = null!;
    
    [JsonIgnore]
    [ValidateNever]
    public DiscordServer Server { get; set; } = null!;
}