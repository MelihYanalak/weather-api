import json
import redis

r = redis.Redis(host='localhost', port=9851)

with open("new_york.geojson", 'r') as file:
    data = json.load(file)

r.execute_command('DROP', 'test_collection')

for idx, feature in enumerate(data['features']):
    id = f"feature_{idx}"
    geojson_str = json.dumps(feature)
    r.execute_command('SET', 'test_collection', id, 'OBJECT', geojson_str)

print("GeoJSON data has been loaded into Tile38")

#dependencies python and redis : pip install redis