using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc.ModelBinding.Validation;

namespace Backend.Models;

public class ServerMember
{
    public string UserId { get; set; }
    
    public string ServerId { get; set; }
    
    [JsonIgnore]
    [ValidateNever]
    public DiscordUser User { get; set; }
    [JsonIgnore]
    [ValidateNever]
    public DiscordServer Server { get; set; }
}

public class AddServerMemberRequest
{
    public string UserId { get; set; }
}