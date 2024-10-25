import requests

from bs4 import BeautifulSoup

url = 'https://quotes.toscrape.com/'

response = requests.get(url)
result = BeautifulSoup(response.text, 'lxml')
quotes = result.find_all('span', class_='text')
for i in quotes :
    print(i.text)