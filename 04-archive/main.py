import os
import time
import shutil
source_dir = ['./source']
target_dir = './target'
if not os.path.exists(target_dir):
    os.mkdir(target_dir)
arcname = target_dir + os.sep + time.strftime('%Y-%m-%d_%H:%M:%S')
try:
    shutil.make_archive(arcname, 'zip', ' '.join(source_dir))
    print(f'Successful backup to {arcname}.zip')
except:
    print('Backup FAILED')
