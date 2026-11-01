using Backend.Data;
using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Services;

public class ServersService
{
    private readonly DiscordDbContext _dbContext;

    public ServersService(DiscordDbContext dbContext)
    {
        _dbContext = dbContext;
    }

    public async Task<DiscordServer> CreateAsync(DiscordServer server)
    {
        var existing = await _dbContext.Servers.FirstOrDefaultAsync(s => s.Id == server.Id);
        if (existing != null)
            return existing;

        await _dbContext.Servers.AddAsync(server);
        await _dbContext.SaveChangesAsync();

        return server;
    }

    public async Task<DiscordServer?> GetByIdAsync(string id)
    {
        return await _dbContext.Servers.FindAsync(id);
    }

    public async Task<bool> DeleteAsync(string id)
    {
        var server = await _dbContext.Servers.FindAsync(id);
        if (server != null)
        {
            _dbContext.Servers.Remove(server);
            await _dbContext.SaveChangesAsync();

            return true;
        }

        return false;
    }
}