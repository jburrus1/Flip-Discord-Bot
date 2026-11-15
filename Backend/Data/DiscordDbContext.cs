using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Data;

public class DiscordDbContext : DbContext
{
    public DiscordDbContext(
        DbContextOptions<DiscordDbContext> options)
        : base(options)
    {
    }

    public DbSet<DiscordUser> Users => Set<DiscordUser>();
    public DbSet<DiscordMessage> Messages => Set<DiscordMessage>();
    public DbSet<DiscordChannel> Channels => Set<DiscordChannel>();
    public DbSet<DiscordServer> Servers => Set<DiscordServer>();
    public DbSet<ServerMember> ServerMembers => Set<ServerMember>();
    

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        // Message -> User
        modelBuilder.Entity<DiscordMessage>()
            .HasOne(m => m.Author)
            .WithMany()
            .HasForeignKey(m => m.AuthorId)
            .OnDelete(DeleteBehavior.Restrict);

        // Message -> Channel
        modelBuilder.Entity<DiscordMessage>()
            .HasOne(m => m.Channel)
            .WithMany()
            .HasForeignKey(m => m.ChannelId)
            .OnDelete(DeleteBehavior.Restrict);

        // Channel -> Server
        modelBuilder.Entity<DiscordChannel>()
            .HasOne(c => c.Server)
            .WithMany()
            .HasForeignKey(c => c.ServerId)
            .OnDelete(DeleteBehavior.Restrict);
        
        // ServerMember CompositeKey
        modelBuilder.Entity<ServerMember>()
            .HasKey(m => new { m.ServerId, m.UserId });
        
        // ServerMember -> User
        modelBuilder.Entity<ServerMember>()
            .HasOne(m => m.User)
            .WithMany()
            .HasForeignKey(m => m.UserId)
            .OnDelete(DeleteBehavior.Restrict);
        
        // ServerMember -> Server
        modelBuilder.Entity<ServerMember>()
            .HasOne(m => m.Server)
            .WithMany()
            .HasForeignKey(m => m.ServerId)
            .OnDelete(DeleteBehavior.Restrict);
    }
}