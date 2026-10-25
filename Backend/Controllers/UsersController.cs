using Backend.Models;
using Backend.Services;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers;

[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    private readonly UsersService _usersService;

    public UsersController(UsersService usersService)
    {
        _usersService = usersService;
    }
    
    [HttpPost]
    public async Task<ActionResult<DiscordUser>> Create(DiscordUser user)
    {
        var createdUser = await _usersService.CreateAsync(user);
        return CreatedAtAction(nameof(GetById), new { id = createdUser.Id }, createdUser);
    }

    [HttpGet]
    public async Task<ActionResult<List<DiscordUser>>> GetAll()
    {
        var users = await _usersService.GetAllAsync();
        return Ok(users);
    }
    
    [HttpGet("{id}")]
    public async Task<ActionResult<DiscordUser>> GetById(string id)
    {
        var user = await _usersService.GetAsync(id);
        if (user == null)
        {
            return NotFound();
        }

        return Ok(user);
    }
    
    [HttpDelete("{id}")]
    public async Task<ActionResult> Delete(string id)
    {
        await _usersService.DeleteAsync(id);
        return NoContent();
    }
    
    [HttpDelete]
    public async Task<ActionResult> DeleteAll()
    {
        await _usersService.DeleteAllAsync();
        return NoContent();
    }
}