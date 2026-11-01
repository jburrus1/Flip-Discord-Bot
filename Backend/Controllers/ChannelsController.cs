using Backend.Models;
using Backend.Services;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers;

[ApiController]
[Route("api/[controller]")]
public class ChannelsController : ControllerBase 
{
    private readonly ChannelsService _channelsService;

    public ChannelsController(ChannelsService channelsService)
    {
        _channelsService = channelsService;
    }
    
    [HttpPost]
    public async Task<ActionResult<DiscordChannel>> Create(DiscordChannel channel)
    {
        await _channelsService.CreateAsync(channel);
        return CreatedAtAction(nameof(GetById), new { id = channel.Id }, channel);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DiscordMessage>> GetById(string id)
    {
        var message = await _channelsService.GetByIdAsync(id);
        if (message == null)
        {
            return NotFound();
        }

        return Ok(message);
    }

    [HttpGet("server/{serverId}")]
    public async Task<ActionResult<List<DiscordMessage>>> GetByServerId(string serverId)
    {
        var messages = await _channelsService.GetByServerIdAsync(serverId);
        return Ok(messages);
    }
    [HttpDelete("{id}")]
    public async Task<ActionResult> Delete(string id)
    {
        await _channelsService.DeleteAsync(id);
        return NoContent();
    }
}