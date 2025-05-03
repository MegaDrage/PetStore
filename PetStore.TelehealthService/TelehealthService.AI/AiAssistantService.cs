using System.Net.Http.Headers;
using System.Text;
using Microsoft.Extensions.Options;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using TelehealthService.Core.Abstractions;

namespace Telehealth.AI;

public class AiAssistantService : IAiAssistantService
{
    private readonly AiServiceSettings _settings;
    private readonly HttpClient _httpClient;
    
    public AiAssistantService(
        IOptions<AiServiceSettings> settings,
        IHttpClientFactory httpClientFactory)
    {
        _settings = settings.Value;
        _httpClient = httpClientFactory.CreateClient("GigaChatClient");
    }

   public async Task<string> GetAccessTokenAsync()
   {
        try
        {
            if (string.IsNullOrEmpty(_settings.ClientSecret))
                throw new ArgumentException("ClientSecret is required");

            var content = new FormUrlEncodedContent(new[]
            {
                new KeyValuePair<string, string>("scope", _settings.Scope),
            });
            
            content.Headers.ContentType = new MediaTypeHeaderValue("application/x-www-form-urlencoded");
            
            var request = new HttpRequestMessage(HttpMethod.Post, _settings.AuthUrl)
            {
                Content = content
            };
            
            request.Headers.Authorization = new AuthenticationHeaderValue("Basic", _settings.ClientSecret);
            request.Headers.Add("RqUID", Guid.NewGuid().ToString());
            request.Headers.Add("Accept", "application/json");
            
            var response = await _httpClient.SendAsync(request);
            var responseText = await response.Content.ReadAsStringAsync();
            
            if (!response.IsSuccessStatusCode)
            {
                throw new HttpRequestException($"HTTP Error {(int)response.StatusCode}: {responseText}");
            }
            
            var responseObject = JObject.Parse(responseText);
            return responseObject["access_token"]?.ToString() 
                ?? throw new Exception("Access token missing in response");
        }
        catch (Exception ex)
        {
            throw new Exception($"GigaChat Auth Failed: {ex.Message}");
        }
   }

    public async Task<string> GetResponseAsync(string prompt)
    {
        var accessToken = await GetAccessTokenAsync();
        
        var request = new HttpRequestMessage(HttpMethod.Post, _settings.ApiUrl);
        request.Headers.Authorization = new AuthenticationHeaderValue("Bearer", accessToken);
        request.Content = new StringContent(
            JsonConvert.SerializeObject(new
            {
                model = "GigaChat",
                messages = new[] { new { role = "user", content = prompt } }
            }),
            Encoding.UTF8,
            "application/json");

        var response = await _httpClient.SendAsync(request);
        var json = await response.Content.ReadAsStringAsync();
        
        return JObject.Parse(json)["choices"]![0]!["message"]!["content"]!.ToString();
    }
}