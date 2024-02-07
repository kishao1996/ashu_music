import moviepy.editor as mp
path = '/Users/bytedance/workspace/ashu_music/music/'
# clip = mp.AudioFileClip(path + '今天必须加班，我用老板的命发誓！.m4a')
# clip.preview()
# clip.close()
# print(clip.nchannels)
# clip = mp.VideoFileClip(path + '今天必须加班，我用老板的命发誓！.m4a')
# clip.audio.write_audiofile(path + 'audio.mp3')
from moviepy import config
exe = config.FFMPEG_BINARY
print(exe)
print(config.try_cmd([config.FFMPEG_BINARY]))