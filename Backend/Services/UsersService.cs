using Backend.Data;
using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Services;

public class UsersService
{
    private readonly DiscordDbContext _dbContext;

    public UsersService(DiscordDbContext dbContext)
    {
        _dbContext = dbContext;
    }

    public async Task<DiscordUser> CreateAsync(DiscordUser user)
    {
        _dbContext.Users.Add(user);
        await _dbContext.SaveChangesAsync();

        return user;
    }
    
    // TODO: update nicknames if we want to store them
    public async Task<bool> UpdateAsync(string id, DiscordUser user)
    {
        var existingUser = await _dbContext.Users.FindAsync(id);
        if (existingUser != null)
        {
            existingUser.Username = user.Username;
            await _dbContext.SaveChangesAsync();
            return true;
        }
        return false;
    }

    public async Task<bool> DeleteAsync(string id)
    {
        var existingUser = await _dbContext.Users.FindAsync(id);
        if (existingUser != null)
        {
            _dbContext.Users.Remove(existingUser);
            await _dbContext.SaveChangesAsync();
            return true;
        }
        return false;
    }
    
    public async Task<bool> DeleteAllAsync()
    {
        _dbContext.Users.RemoveRange(_dbContext.Users);
        await _dbContext.SaveChangesAsync();
        return true;
    }
    
    public async Task<DiscordUser?> GetAsync(string id)
    {
        return await _dbContext.Users.FindAsync(id);
    }
    
    public async Task<List<DiscordUser>> GetAllAsync()
    {
        return await _dbContext.Users.ToListAsync();
    }
}