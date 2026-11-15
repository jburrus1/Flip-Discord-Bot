namespace Frontend.State;

public class ServerSelectionState
{
    public event Action? OnChange;
    public string? SelectedServerId { get; private set; }

    public void SelectServer(string serverId)
    {
        if (SelectedServerId == serverId)
            return;

        SelectedServerId = serverId;
        OnChange?.Invoke();
    }
}