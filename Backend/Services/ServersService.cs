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

    public async Task<List<DiscordServer>> GetAllAsync()
    {
        return await _dbContext.Servers.ToListAsync();
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

    public async Task<ServerMember?> AddUserToServerAsync(string serverId, string userId)
    {
        var existing = await _dbContext.ServerMembers.FirstOrDefaultAsync(m => m.ServerId == serverId && m.UserId == userId);
        if (existing != null)
        {
            return existing;
        }

        var member = new ServerMember
        {
            ServerId = serverId,
            UserId = userId
        };
        

        await _dbContext.ServerMembers.AddAsync(member);
        await _dbContext.SaveChangesAsync();

        return member;
    }

    public async Task<List<DiscordUser>> GetUsersAsync(string serverId)
    {
        return await _dbContext.ServerMembers
            .Where(m => m.ServerId == serverId)
            .Select(m => m.User).ToListAsync();
    }

    public async Task<List<DiscordMessage>> GetMessagesAsync(string serverId)
    {
        return await _dbContext.Messages.Where(m => m.Channel.ServerId == serverId).ToListAsync();
    }
    public async Task<bool> ExistsAsync(string serverId)
    {
        return await _dbContext.Servers
            .AnyAsync(s => s.Id == serverId);
    }
}